var UIMAGE = {};

UIMAGE.roi = function(img, x, y, w, h) {
	var col = img.width,
	    row = img.height,
		c = 4;
    var area = w * h;
	var x1 = x-Math.ceil(w/2),
		y1 = y-Math.ceil(h/2);
    var roiBuff = new Uint8ClampedArray(area * c);
	for(var i=0;i<h;i++){
		var t = 4*(y1-1+i)*col+4*x1;
	    var tr = 4*i*w;
	    var tLength = 4*w;
		for(var j=0;j<tLength;j++)
			roiBuff[tr+j] = img.data[t+j];
	}
	
    return roiBuff;
}

UIMAGE.addWeighted = function(patchBuff, roiBuff, p1, p2,w,h){
	var area = w*h,
	    c = 4;
	var dstBuff = new Uint8ClampedArray(area * c);
	for(var i=0;i<area*c;i++){
		var pb = patchBuff[i];
		var rb = roiBuff[i]
		dstBuff[i] = rb*p2 +pb*p1;
	}
	return dstBuff;
}

UIMAGE.pickX = function(src){
	var x = src.data[0]*100.0+src.data[4];
	return x
}

UIMAGE.pickY = function(src){
	var y = src.data[8]*100.0+src.data[12];
	return y
}

UIMAGE.imgCopy = function(src){
  var w = src.width,
      h = src.height;
	var c = 4;
	var area = w * h;
	var dst = new Uint8ClampedArray(area * c);
	for(var i=0;i<area*c;i++){
		dst[i] = src.data[i];
	}
	return dst
}

UIMAGE.roiAdd = function(img, roiBuff, x, y,w,h) {
	var col = img.width,
	    row = img.height,
		c = 4;
    var area = w * h;
	var x1 = x-Math.ceil(w/2),
		y1 = y-Math.ceil(h/2);
	for(var i=0;i<h;i++){
		var t = c*(y1-1+i)*col+c*x1;
	    var tr = c*i*w;
	    var tLength = c*w;
		for(var j=0;j<tLength;j++)
			img.data[t+j] = roiBuff[tr+j];
	}
	return img
}

UIMAGE.bgr2inten = function(imgBuff,w,h) {
	var intensity = new Array(h);
	for (var y=0;y<h;y++){
		intensity[y] = new Array(w);
		
		for (var x=0;x<w;x++){
			var i = x*4+y*4*w;
			var r = imgBuff[i],
			g = imgBuff[i+1],
			b = imgBuff[i+2],
			a = imgBuff[i+3];
			
	    var luma =(r * 299/1000 + g * 587/1000+ b * 114/1000) / 255;

	    intensity[y][x] = luma;
		}
	}
	return intensity;
}

UIMAGE.similarity = function(des1,des2) {
	var n,x1,x2,sim,sum;
	n = des1.length;
	sum = 0;
	x1 = 0;
	x2 = 0;

	for(var i=0;i<n;i++){
		x1 = x1+des1[i]*des1[i];
		x2 = x2+des2[i]*des2[i];
		sum = sum+des1[i]*des2[i];
	}
	x1 = Math.sqrt(x1);
	x2 = Math.sqrt(x2);
	sim = sum/(x1*x2)
	return sim;
}

UIMAGE.arrayToImage = function(buffArray,img) {
	var w = img.width,
		h = img.height,
		c = 4;
	var length = w*h*c;
	for(var i=0;i<length;i++)
		img.data[i] = buffArray[i]
	return img
}


module.exports = UIMAGE
