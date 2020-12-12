import * as files from 'fs';

let input = files.readFileSync('12.txt').toString().split('\n').map((a) => a.split(new RegExp('\(\[A-Z\]\+\)\(\[0-9\]\+\)')))
	.map((a) => a.slice(1, 3));

let NS = 0;
let EW = 0;
let angle = 0;

for (let i = 0; i < input.length; i++) {
	if (input[i][0] === 'N') {
		NS += Number(input[i][1]);
	} else if (input[i][0] === 'S') {
		NS -= Number(input[i][1]);
	} else if (input[i][0] === 'E') {
		EW += Number(input[i][1]);
	} else if (input[i][0] === 'W') {
		EW -= Number(input[i][1]);
	} else if (input[i][0] === 'F') {
		if (angle === 0) {
			EW += Number(input[i][1]);
		} else if (angle === 90) {
			NS += Number(input[i][1]);
		} else if (angle === 180) {
			EW -= Number(input[i][1]);
		} else if (angle === 270) {
			NS -= Number(input[i][1]);
		}
	} else if (input[i][0] === 'L') {
		console.log("Left", input[i][1]);

		rotateLeft(Number(input[i][1]));
	} else if (input[i][0] === 'R') {
		console.log("Right", input[i][1]);
		rotateRight(Number(input[i][1]));
	}
}

function rotateLeft(increment: number){
	angle = (angle + increment) % 360;
	console.log(angle);

}

function rotateRight(increment: number){
	if (angle - increment < 0) {
		angle += (360 - increment);
	} else {
		angle -= increment;
	}
	console.log(angle);

}

console.log("1st Part: ", Math.abs(NS) + Math.abs(EW));


