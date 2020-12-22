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
		rotateLeft(Number(input[i][1]));
	} else if (input[i][0] === 'R') {
		rotateRight(Number(input[i][1]));
	}
}

function rotateLeft(increment: number) {
	angle = (angle + increment) % 360;
}

function rotateRight(increment: number) {
	if (angle - increment < 0) {
		angle += (360 - increment);
	} else {
		angle -= increment;
	}
}

console.log("1st Part: ", Math.abs(NS) + Math.abs(EW));


let waypoint = [1, 10];
let position = [0, 0];

for (let i = 0; i < input.length; i++) {
	if (input[i][0] === 'N') {
		waypoint[0] += Number(input[i][1]);
	} else if (input[i][0] === 'S') {
		waypoint[0] -= Number(input[i][1]);
	} else if (input[i][0] === 'E') {
		waypoint[1] += Number(input[i][1]);
	} else if (input[i][0] === 'W') {
		waypoint[1] -= Number(input[i][1]);
	} else if (input[i][0] === 'F') {
		position[0] += Number(input[i][1]) * waypoint[0];
		position[1] += Number(input[i][1]) * waypoint[1];
	} else if (input[i][0] === 'L') {
		rotateLeft2(Number(input[i][1]));
	} else if (input[i][0] === 'R') {
		rotateRight2(Number(input[i][1]));
	}
	console.log(input[i], waypoint, position);

}

function rotateLeft2(angle: number) {
	for (let i = 0; i < angle / 90; i++){
		let temp = waypoint[1];
		waypoint[1] = -waypoint[0];
		waypoint[0] = temp;
	}
}

function rotateRight2(angle: number) {
	for (let i = 0; i < angle / 90; i++){
		let temp = waypoint[1];
		waypoint[1] = waypoint[0];
		waypoint[0] = -temp;
	}
}

console.log('2nd Part: ', Math.abs(position[0]) + Math.abs(position[1]));
