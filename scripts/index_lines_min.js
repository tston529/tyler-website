function sleep(ms){return new Promise(resolve=>setTimeout(resolve,ms))}var c=document.getElementById("hero");var ctx=c.getContext("2d");var m=[20].fill(0).map(function(n){return Math.floor(Math.random()*20)+1});var xCoords=[20].fill(0).map(function(n){return Math.floor(Math.random()*(1024-10)+10)});var yCoords=[20].fill(0).map(function(n){return Math.floor(Math.random()*(768-10)+10)});async function drawLines(ctx,m,xCoords,yCoords){ctx.imageSmoothingEnabled=true;ctx.strokeStyle="#c4b56c";ctx.beginPath();ctx.moveTo(xCoords[0],yCoords[0]);ctx.arc(xCoords[0],yCoords[0],5,0,2*Math.PI);ctx.stroke();for(var i=1;i<20;i+=1){ctx.lineTo(xCoords[m[i]],yCoords[m[i]]);ctx.stroke();ctx.arc(xCoords[m[i]],yCoords[m[i]],5,0,2*Math.PI);ctx.stroke();ctx.moveTo(xCoords[m[i]],yCoords[m[i]]);}ctx.lineTo(xCoords[0],yCoords[0]);ctx.stroke();ctx.arc(xCoords[0],yCoords[0],5,0,2*Math.PI);ctx.stroke();ctx.closePath()}async function main(c,ctx,m,xCoords,yCoords){var frames=100;var accelSpeed=(frames/frames)*3;var i=0;while(accelSpeed>0){accelSpeed=((frames-i)/frames)*2;ctx.clearRect(0,0,c.width,c.height);drawLines(ctx,m,xCoords,yCoords);for(var j=0;j<20;j+=1){if(xCoords[j]>10&&xCoords[j]<c.width/2){xCoords[j]-=Math.random()*accelSpeed+0.5}else if(xCoords[j]<c.width-20&&xCoords[j]>c.width/2){xCoords[j]+=Math.random()*accelSpeed+0.5}if(yCoords[j]>20&&yCoords[j]<c.height/2){yCoords[j]-=Math.random()*accelSpeed+0.5}else if(yCoords[j]<c.height-20&&yCoords[j]>c.height/2){yCoords[j]+=Math.random()*accelSpeed+0.5}}await sleep(50);i+=1}}main(c,ctx,m,xCoords,yCoords);