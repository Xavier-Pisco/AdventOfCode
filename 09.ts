import * as files from 'fs';

let input = files.readFileSync('09.txt').toString().split('\n').map((a) => Number(a));

function checkSum(previous25: number[], num: number) {
	for (let i = 0; i < previous25.length - 1; i++){
		for (let j = 0; j < previous25.length; j++){
			if (previous25[i] + previous25[j] === num) return true;
		}
	}
	return false;
}

let previous25: number[] = [];

for (let i = 0; i < 25; i++){
	previous25.push(input[i]);
}

for (let i = 25; i < input.length; i++){
	if (!checkSum(previous25, input[i])) {
		console.log('1st Part: ' + input[i]);
		break;
	}
	previous25.shift();
	previous25.push(input[i]);
}

previous25 = [];

for (let i = 0; i < 25; i++){
	previous25.push(input[i]);
}

for (let i = 25; i < input.length; i++){
	if (!checkSum(previous25, input[i])) {
		let sum: number = 0;
		for (let j = 0; j < i - 1; j++){
			let list: number[] = [];
			for (let x = j; x < i; x++){
				list.push(input[x]);
				sum = 0;
				for (let y = 0; y < list.length; y++){
					sum += list[y];
				}
				if (sum === input[i]){
					list = list.sort((a,b) => a - b);
					console.log('2nd Part: ' + (list[0] + list[list.length - 1]));
					break;
				}
			}
			if (sum === input[i]){
				break;
			}
		}
		break;
	}
	previous25.shift();
	previous25.push(input[i]);

}