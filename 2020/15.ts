import * as files from 'fs';

let numbers: number[] = [];
let i = 0;
let previous: number;

let input = files.readFileSync('15.txt').toString().split(',').map((a) => {
	previous !== undefined ? numbers[previous] = i : i;
	previous = Number(a);
	i++;
});

for (i; i < 2020; i++) {
	if (numbers[previous] !== undefined) {
		let temp = i - numbers[previous];
		numbers[previous] = i;
		previous = temp;
	} else {
		numbers[previous] = i;
		previous = 0;
	}

}

console.log('1st Part: ', previous);

numbers = [];
i = 0;
previous = undefined;

input = files.readFileSync('15.txt').toString().split(',').map((a) => {
	previous !== undefined ? numbers[previous] = i : i;
	previous = Number(a);
	i++;
});

for (i; i < 30000000; i++) {
	if (numbers[previous] !== undefined) {
		let temp = i - numbers[previous];
		numbers[previous] = i;
		previous = temp;
	} else {
		numbers[previous] = i;
		previous = 0;
	}
}

console.log('2nd Part: ', previous);
