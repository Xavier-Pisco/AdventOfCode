import * as files from 'fs';

let input: string[] = files.readFileSync('03.txt').toString().split('\n');

function treesEncounters(xincrement: number, yincrement: number) {
	let x = 0;
	let trees = 0;
	for (let y = 0; y < input.length; y += yincrement) {
		if (input[y][x] == '#') trees += 1;
		x = (x + xincrement) % input[y].length;
	}
	return trees;
}

console.log('1st Part: ' + treesEncounters(3, 1));

console.log('2nd Part: ' +
	(treesEncounters(1, 1) * treesEncounters(3, 1) *
	treesEncounters(5, 1) * treesEncounters(7, 1) * treesEncounters(1, 2)));


