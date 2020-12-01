import * as files from 'fs';


let input : number[] = files.readFileSync('01.txt').toString().split('\n').map(Number);

for (let i : number = 0; i < input.length - 1; i++){
	for (let j : number = i; j < input.length; j++){
		if (input[i] + input[j] === 2020) {
			console.log(input[i] + ' * '  + input[j] + ' = ' + input[i] * input[j]);
		}
	}
}
