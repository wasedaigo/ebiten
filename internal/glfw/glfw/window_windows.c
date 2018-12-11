#include "_cgo_export.h"

void glfwWindowPosCB(GLFWwindow* window, int xpos, int ypos) {
	goWindowPosCB(window, xpos, ypos);
}

void glfwWindowSizeCB(GLFWwindow* window, int width, int height) {
	goWindowSizeCB(window, width, height);
}

void glfwFramebufferSizeCB(GLFWwindow* window, int width, int height) {
	goFramebufferSizeCB(window, width, height);
}

void glfwWindowCloseCB(GLFWwindow* window) {
	goWindowCloseCB(window);
}

void glfwWindowFocusCB(GLFWwindow* window, int focused) {
	goWindowFocusCB(window, focused);
}

void glfwSetWindowPosCallbackCB(GLFWwindow* window) {
	glfwSetWindowPosCallback(window, glfwWindowPosCB);
}

void glfwSetWindowSizeCallbackCB(GLFWwindow* window) {
	glfwSetWindowSizeCallback(window, glfwWindowSizeCB);
}

void glfwSetFramebufferSizeCallbackCB(GLFWwindow* window) {
	glfwSetFramebufferSizeCallback(window, glfwFramebufferSizeCB);
}

void glfwSetWindowCloseCallbackCB(GLFWwindow* window) {
	glfwSetWindowCloseCallback(window, glfwWindowCloseCB);
}

void glfwSetWindowFocusCallbackCB(GLFWwindow* window) {
	glfwSetWindowFocusCallback(window, glfwWindowFocusCB);
}
