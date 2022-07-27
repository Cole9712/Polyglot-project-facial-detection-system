# Facial Detection Web System

This is a web system that supports keyword search images, facial detection and face swapping

## Topic Idea

If we want to perform some advanced/complex image processing on the web only using JavaScript, it is really hard. To solve this, the project creates a web system that can perform **facial detection** and **face swapping** from user-uploaded images.
For image processing, I implemented it with **Python** since it has a large number of supported libraries. Moreover, it is easy to handle requests concurrently using **Go** at the backend.\
I also implemented facial detection using image API from [Pixabay. com](https://pixabay.com/), users can search images by keyword and detect faces on those images.

## Languages

### Front-end

- JavaScript with [Vue](https://vuejs.org/)

### Back-end

- Go
- Python with OpenCV

## Inter-language communication methods

- Command line execution (Go &leftarrow;&rightarrow; Python)
- RPC server (JS &leftarrow;&rightarrow; Go)

## Deployment technology

Vagrant

## Installation

**Prerequisite:** Vagrant and VirtualBox.

The first build time might be as long as **20 minutes** due to the long compiling time for 'dlib' package, which is an essential library for facial landmark detection.

**If you are deploying with Windows, make sure run Vagrant commands as Admin since Windows only allows admin to create symbolic links.**

`vagrant up` This command creates and configures guest machines according to your Vagrantfile.

After successful installation, access `http://localhost:5555/` on web browser to use front-end client.

## Features (with demo)

The system can accept jpeg/jpg/png format image files.

1. Home Page: search images via keywords, you can enlarge them or detect faces on them.

![Api Demo](md_assets/api_demo.gif)

2. Facial Detect Page: upload your own image to perform facial detection. Note that due to the limitation of algorithm(not a AI engine after all), there might be some faces that could not be detected.*There are some sample images available to test in `test_imgs` directory.*

![Detection Demo](md_assets/detection_demo.gif)

3. Face Swapping Page: upload two images(say image 1 and image 2), a confirmation alert will pop up and you can swap image 2's face onto image 1's person. To maximize the face swapping result quality, please upload two images that share close resolution value.*There are some sample images available to test in the `test_imgs` directory under the main repository.*

![Swapping Demo](md_assets/face_swap_demo.gif)

*All Pages can be accessed via top navigation bar.

## Dependencies

- [v-uploader](https://github.com/TerryZ/v-uploader)
- [v-dialogs](https://github.com/TerryZ/v-dialogs)
- [Vue Material](https://vuematerial.io/)
