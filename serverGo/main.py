import cv2
import numpy as np
import pkg_resources
import sys
import dlib

# CONSTANTS
haar_xml = pkg_resources.resource_filename(
    'cv2', 'data/haarcascade_frontalface_default.xml')
predictor = dlib.shape_predictor('data/shape_predictor_68_face_landmarks.dat')

RIGHT_EYE_POINTS = range(36, 42)
LEFT_EYE_POINTS = range(42, 48)

LEFT_BROW_POINTS = range(22, 27)
RIGHT_BROW_POINTS = range(17, 22)
NOSE_POINTS = range(27, 35)
MOUTH_POINTS = range(48, 61)

POINTS = [RIGHT_EYE_POINTS, LEFT_EYE_POINTS, LEFT_BROW_POINTS, RIGHT_BROW_POINTS, NOSE_POINTS, MOUTH_POINTS]
FLAT_POINTS = range(68)
ALIGN_POINTS = (LEFT_BROW_POINTS + RIGHT_EYE_POINTS + LEFT_EYE_POINTS +
                               RIGHT_BROW_POINTS + NOSE_POINTS + MOUTH_POINTS)

def convertToBGR(image):
    return cv2.cvtColor(image, cv2.COLOR_RGB2BGR)


def convertToRGB(image):
    return cv2.cvtColor(image, cv2.COLOR_BGR2RGB)


def convertToGRAY(image):
    return cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)


def face_detect(img):
    front_face_cas = cv2.CascadeClassifier(haar_xml)
    if img is None:
        print("Image is None")
        return None
    face_img = img.copy()
    rect = front_face_cas.detectMultiScale(
        face_img, scaleFactor=1.2, minNeighbors=5)
    for (x, y, w, h) in rect:
        cv2.rectangle(face_img, (x, y), (x+w, y+h), (255, 255, 255), 4)
    return face_img

# Face swapping algorithm adapted from http://matthewearl.github.io/2015/07/28/switching-eds-with-python/
# modified to use haar cascade
def get_landmarks(img, using_dlib = True):
    cascade = cv2.CascadeClassifier(haar_xml)
    dlib_detector = dlib.get_frontal_face_detector()
    rect = dlib_detector(img, 1) if using_dlib else cascade.detectMultiScale(
        img, scaleFactor=1.05, minNeighbors=5)
    
    # if len(rect) != 1:
    #     raise RuntimeError("Nr. of faces: "+str(len(rect)))
    if len(rect) == 0:
        raise RuntimeError("Nr. of faces: "+str(len(rect)))
    
    if using_dlib:
        single_rect = rect[0]
    else:
        x, y, w, h=rect[0]
        single_rect = dlib.rectangle(int(x),int(y),int(x+w),int(y+h))
    return np.matrix([[p.x, p.y] for p in predictor(img, single_rect).parts()])

def get_face_mask(img, landmarks):
    img = np.zeros(img.shape[:2], dtype=np.float64)
    
    for group in POINTS:
        points = cv2.convexHull(landmarks[group])
        cv2.fillConvexPoly(img, points, color=1)

    img = np.array([img, img, img]).transpose((1, 2, 0))
    FEATHER_AMOUNT = int(11)
    # FEATHER_AMOUNT = int(18.75)
    img = (cv2.GaussianBlur(img, (FEATHER_AMOUNT, FEATHER_AMOUNT), 0) > 0) * 1.0
    img = cv2.GaussianBlur(img, (FEATHER_AMOUNT, FEATHER_AMOUNT), 0)
    return img

# algorithm from https://en.wikipedia.org/wiki/Orthogonal_Procrustes_problem
def transformation_from_points(points1, points2):
    c1 = np.mean(points1, axis=0)
    c2 = np.mean(points2, axis=0)
    s1 = np.std(points1)
    s2 = np.std(points2)

    points1 = (np.array(points1)-c1)/s1
    points2 = (np.array(points2)-c2)/s2
    
    U, S, Vt = np.linalg.svd(points1.T * points2)

    R = (U * Vt).T
    return np.vstack([  np.hstack(((s2 / s1) * R,
                    c2.T - (s2 / s1) * R * c1.T)),
                    np.matrix([0., 0., 1.])])

def warp_im(img, M, dshape):
    output_im = np.zeros(dshape, dtype=img.dtype)
    cv2.warpAffine(img,
                   M[:2],
                   (dshape[1], dshape[0]),
                   dst=output_im,
                   borderMode=cv2.BORDER_TRANSPARENT,
                   flags=cv2.WARP_INVERSE_MAP)
    return output_im

def correct_colours(im1, im2, landmarks1):
    blur_amount = 0.6 * np.linalg.norm(
                              np.mean(landmarks1[LEFT_EYE_POINTS], axis=0) -
                              np.mean(landmarks1[RIGHT_EYE_POINTS], axis=0))
    blur_amount = int(blur_amount)//2*2+1
    im1_blur = cv2.GaussianBlur(im1, (blur_amount, blur_amount), 0)
    im2_blur = cv2.GaussianBlur(im2, (blur_amount, blur_amount), 0)

    im2_blur += (128 * (im2_blur <= 1.0)).astype(im2_blur.dtype)

    return (im2.astype(np.float64) * im1_blur.astype(np.float64) / im2_blur.astype(np.float64))

def face_swap(img1, img2):
    SCALE_FACTOR = 1
    img1 = cv2.resize(img1,None,fx=1, fy=1, interpolation = cv2.INTER_LINEAR)
    img1 = cv2.resize(img1, (img1.shape[1] * SCALE_FACTOR, img1.shape[0] * SCALE_FACTOR))
    img1LM = get_landmarks(img1)

    # img2 = cv2.resize(img2, None, fx=0.75, fy=0.75, interpolation=cv2.INTER_LINEAR)

    img2LM = get_landmarks(img2)
    img2mask = get_face_mask(img2,img2LM)
    img2 = cv2.resize(img2,None,fx=1, fy=1, interpolation = cv2.INTER_LINEAR)
    img2 = cv2.resize(img2, (img2.shape[1] * SCALE_FACTOR, img2.shape[0] * SCALE_FACTOR))
    #  get tranformation matrix
    M = transformation_from_points(img1LM[FLAT_POINTS], img2LM[FLAT_POINTS])

    warped_mask = warp_im(img2mask, M, img1.shape)
    warped_img2 = warp_im(img2, M, img1.shape)

    all_in_one_mask = np.max([get_face_mask(img1, img1LM), warped_mask],axis=0)

    # perform extra color correction
    warped_img2 = correct_colours(img1, warped_img2, img1LM)

    result = warped_img2 * all_in_one_mask + img1 * (1.0 - all_in_one_mask)

    return result



if __name__ == "__main__":
    numOfArgvs = len(sys.argv) - 1
    if numOfArgvs == 0 or numOfArgvs == 1:
        sys.exit("Too few arguments")
    execType = sys.argv[1]
    if execType=="detect":
        fileName = sys.argv[2]
        pic = cv2.imread('output/' + fileName, cv2.IMREAD_COLOR)
        result = face_detect(pic)
        cv2.imwrite("facialOutput/" + fileName, result)
    elif execType == "swap" and numOfArgvs == 3:
        fileName1 = sys.argv[2]
        fileName2 = sys.argv[3]
        pic1 = cv2.imread('output/' + fileName1, cv2.IMREAD_COLOR)
        pic2 = cv2.imread('output/' + fileName2, cv2.IMREAD_COLOR)
        try:
            pic1 = face_swap(pic1, pic2)
        except RuntimeError as e:
            pic1 = cv2.putText(pic1, str(e), (5, 20), cv2.FONT_HERSHEY_SIMPLEX, 0.5,(255,255,255),1,cv2.LINE_AA)
        cv2.imwrite("facialOutput/" + fileName1, pic1)
    else :
        print("Arguments not meet format requirements.")
        
    # pic = convertToBGR(pic) 

    cv2.destroyAllWindows()
