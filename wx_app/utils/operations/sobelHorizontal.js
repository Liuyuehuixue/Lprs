import convolution from './convolution'

const sobelHorizontal = function (pixels, width, height) {
  const divider = 4;
  const kernel = [1 / divider, 2 / divider, 1 / divider, 0, 0, 0, -1 / divider, -2 / divider, -1 / divider];

  pixels = convolution(pixels, kernel, width, height);
  return pixels
}

export default sobelHorizontal
