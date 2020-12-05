import { LOADIPHLPAPI } from 'dns';
import * as files from 'fs';

let input: string[] = files.readFileSync('05.txt').toString().split('\n');

let rows: number[] = [];
for (let i = 0; i < 128; i++) {
	rows.push(i);
}


let columns: number[] = [];
for (let i = 0; i < 8; i++) {
	columns.push(i);
}

let highest = 0;

let currentRows: number[];
let currentColumns: number[];
for (let i = 0; i < input.length; i++) {
	currentRows = Array.from(rows);

	currentColumns = Array.from(columns);
	let j: number;
	for (j = 0; j < 7; j++) {
		if (input[i][j] === 'F') {
			currentRows = currentRows.splice(0, currentRows.length / 2);

		} else {
			currentRows = currentRows.splice(currentRows.length / 2, currentRows.length);

		}
	}
	for (j; j < 10; j++) {
		if (input[i][j] === 'R') {
			currentColumns = currentColumns.splice(currentColumns.length / 2, currentColumns.length);
		} else {
			currentColumns = currentColumns.splice(0, currentColumns.length / 2);
		}
	}
	let seat = currentRows[0] * 8 + currentColumns[0];
	if (seat > highest){
		highest = seat;
	}
}

console.log('1st Part: ' + highest);

let seats: number[] = [];

for (let i = 0; i < input.length; i++) {
	currentRows = Array.from(rows);

	currentColumns = Array.from(columns);
	let j: number;
	for (j = 0; j < 7; j++) {
		if (input[i][j] === 'F') {
			currentRows = currentRows.splice(0, currentRows.length / 2);

		} else {
			currentRows = currentRows.splice(currentRows.length / 2, currentRows.length);

		}
	}
	for (j; j < 10; j++) {
		if (input[i][j] === 'R') {
			currentColumns = currentColumns.splice(currentColumns.length / 2, currentColumns.length);
		} else {
			currentColumns = currentColumns.splice(0, currentColumns.length / 2);
		}
	}
	seats.push(currentRows[0] * 8 + currentColumns[0]);
}

seats.sort();

for (let i = 0; i < seats.length - 1; i++){
	if (seats[i] + 1 !== seats[i+1] && seats[i] + 2 === seats[i+1]){
		console.log('2nd Part: ', seats[i] + 1);
		break;
	}
}