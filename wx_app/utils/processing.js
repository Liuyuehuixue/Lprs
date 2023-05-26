var processing = {
  gradients: function(intensities) {
    return this._gradients(intensities);
  },

  _gradients: function(intensities) {
    var height = intensities.length;
    var width = intensities[0].length;

    var gradX = new Array(height);
    var gradY = new Array(height);

    for (var y = 0; y < height; y++) {
      gradX[y] = new Array(width);
      gradY[y] = new Array(width);

      var row = intensities[y];

      for (var x = 0; x < width; x++) {
        var prevX = x == 0 ? 0 : intensities[y][x - 1];
        var nextX = x == width - 1 ? 0 : intensities[y][x + 1];
        var prevY = y == 0 ? 0 : intensities[y - 1][x];
        var nextY = y == height - 1 ? 0 : intensities[y + 1][x];

        // kernel [-1, 0, 1]
        gradX[y][x] = -prevX + nextX;
        gradY[y][x] = -prevY + nextY;
      }
    }

    return {
      x: gradX,
      y: gradY
    };
  },

  gradientVectors: function(intensities) {
    return this._gradientVectors(intensities);
  },

  _gradientVectors: function(intensities) {
    var height = intensities.length;
    var width = intensities[0].length;

    var vectors = new Array(height);

    for (var y = 0; y < height; y++) {
      vectors[y] = new Array(width);

      for (var x = 0; x < width; x++) {
        var prevX = x == 0 ? 0 : intensities[y][x - 1];
        var nextX = x == width - 1 ? 0 : intensities[y][x + 1];
        var prevY = y == 0 ? 0 : intensities[y - 1][x];
        var nextY = y == height - 1 ? 0 : intensities[y + 1][x];

        // kernel [-1, 0, 1]
        var gradX = -prevX + nextX;
        var gradY = -prevY + nextY;

        vectors[y][x] = {
          mag: Math.sqrt(Math.pow(gradX, 2) + Math.pow(gradY, 2)),
          orient: Math.atan2(gradY, gradX)
        }
      }
    }
    return vectors;
  }
}

module.exports = processing;
