import cv2
import numpy as np
import pkg_resources
import sys


def convertToBGR(image):
    return cv2.cvtColor(image, cv2.COLOR_RGB2BGR)


def convertToRGB(image):
    return cv2.cvtColor(image, cv2.COLOR_BGR2RGB)


def convertToGRAY(image):
    return cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)


def face_detect(img):
    haar_xml = pkg_resources.resource_filename(
        'cv2', 'data/haarcascade_frontalface_default.xml')
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


if __name__ == "__main__":
    numOfArgvs = len(sys.argv) - 1
    if numOfArgvs == 0:
        sys.exit("Too few arguments")
    fileName = sys.argv[1]
    pic = cv2.imread('../serverGo/output/' + fileName, cv2.IMREAD_COLOR)

    # pic = convertToBGR(pic)
    result = face_detect(pic)

    cv2.imwrite("output/" + fileName, result)
    cv2.destroyAllWindows()
