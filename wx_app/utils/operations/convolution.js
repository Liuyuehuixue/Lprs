const convolution = function (src, kernel,width,height) {
  let side = Math.round(Math.sqrt(kernel.length))
  let halfSide = Math.floor(side / 2)
  var area = width * height;
  var outputData = new Uint8Array(area * 4);

  for (let y = 0; y < height; y++) {
    for (let x = 0; x < width; x++) {
      let dstOff = (y * width + x) * 4,
        sumReds = 0,
        sumGreens = 0,
        sumBlues = 0

      for (let kernelY = 0; kernelY < side; kernelY++) {
        for (let kernelX = 0; kernelX < side; kernelX++) {
          let currentKernelY = y + kernelY - halfSide,
            currentKernelX = x + kernelX - halfSide

          if (
            currentKernelY >= 0 &&
            currentKernelY < height &&
            currentKernelX >= 0 &&
            currentKernelX < width
          ) {
            let offset = (currentKernelY * width + currentKernelX) * 4,
              weight = kernel[kernelY * side + kernelX]

            sumReds += src[offset] * weight
            sumGreens += src[offset + 1] * weight
            sumBlues += src[offset + 2] * weight
          }
        }
      }

      outputData[dstOff] = sumReds
      outputData[dstOff + 1] = sumGreens
      outputData[dstOff + 2] = sumBlues
      outputData[dstOff + 3] = 255
    }
  }
  return outputData
}

export default convolution
