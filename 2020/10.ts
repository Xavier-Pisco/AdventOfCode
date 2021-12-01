import * as files from 'fs';

let input = files.readFileSync('10.txt').toString().split('\n').map((a) => Number(a));
input.push(0);
input.sort((a, b) => a - b);
input.push(input[input.length - 1] + 3);

let jump1 = 0;
let jump3 = 0;
for (let i = 0; i < input.length - 1; i++) {
	if (input[i + 1] - input[i] === 1) {
		jump1 += 1;
	}
	else if (input[i + 1] - input[i] === 3) {
		jump3 += 1;
	}
}

console.log('1st Part: ', jump1 * jump3);

function countSolutions(input: number[]) {
	let result = 0;


	if (input.length > 3) {
		if (input[1] - input[0] <= 3) {
			result += countSolutions(input.slice(1));
		}
		if (input[2] - input[0] <= 3) {
			result += countSolutions(input.slice(2));
		}
		if (input[3] - input[0] <= 3) {
			result += countSolutions(input.slice(3));
		}
	} else if (input.length > 2) {
		if (input[1] - input[0] <= 3) {
			result += countSolutions(input.slice(1));
		}
		if (input[2] - input[0] <= 3) {
			result += 1;
		}
	} else if (input.length > 1) {
		if (input[1] - input[0] <= 3) {
			result += 1;
		}
	} else {
		return 1;
	}

	return result;
}

let values: number[] = [];
values[input.length - 1] = 1;

for (let i = input.length - 2; i >= 0; i--) {
	if (input[i + 1] - input[i] <= 3 && input[i + 2] - input[i] <= 3 && input[i + 3] - input[i] <= 3) {
		values[i] = values[i + 1] + values[i + 2] + values[i + 3];
	} else if (input[i + 1] - input[i] <= 3 && input[i + 2] - input[i] <= 3) {
		values[i] = values[i + 1] + values[i + 2];
	} else if (input[i + 1] - input[i] <= 3) {
		values[i] = values[i + 1];
	}
}


console.log('2nd Part: ', values[0]);
