import * as files from 'fs';

let input = files.readFileSync('16.txt').toString().split(new RegExp('\n*your ticket:\n*\|\n*nearby tickets:\n*'));

let classes = input[0].split(new RegExp('\n?[a-z]*: '));
classes.splice(0, 1);

let valids: boolean[] = [];
for (let i = 0; i < classes.length; i++) {
	let values = classes[i].split(' or ');
	let values1 = values[0].split('-');
	let values2 = values[1].split('-');
	for (let j = Number(values1[0]); j <= Number(values1[1]); j++) {
		valids[j] = true;
	}
	for (let j = Number(values2[0]); j <= Number(values2[1]); j++) {
		valids[j] = true;
	}
}

let final = 0;
let nearby = input[2].split('\n').map((a) => a.split(','));
for (let i = 0; i < nearby.length; i++) {
	for (let j = 0; j < nearby[i].length; j++) {
		if (valids[Number(nearby[i][j])] !== true)
			final += Number(nearby[i][j]);
	}
}

console.log('1st Part: ', final);

let validTickets = [];

nearby = input[2].split('\n').map((a) => a.split(','));
for (let i = 0; i < nearby.length; i++) {
	for (let j = 0; j < nearby[i].length; j++) {
		if (valids[Number(nearby[i][j])] !== true)
			break
		if (j === nearby[i].length - 1)
			validTickets.push(nearby[i]);
	}
}

let order = [];
for (let i = 0; i < classes.length; i++) {
	let temp = []
	let values = classes[i].split(' or ');
	let values1 = values[0].split('-');
	let values2 = values[1].split('-');
	for (let j = Number(values1[0]); j <= Number(values1[1]); j++) {
		temp[j] = true;
	}
	for (let j = Number(values2[0]); j <= Number(values2[1]); j++) {
		temp[j] = true;
	}
	for (let j = 0; j < validTickets[0].length; j++) {
		if (order[i] !== undefined) break;
		if (contains(order, j)) continue;
		for (let x = 0; x < validTickets.length; x++) {
			if (temp[Number(validTickets[x][j])] !== true)
				break;
			if (x === validTickets.length - 1)
				order[i] = j;
		}
	}
}

function contains(array: any[], value: any) {
	for (let i = 0; i < array.length; i++) {
		if (array[i] === value && value !== undefined)
			return true;
	}

	return false;
}
final = 1;

let yourTickets = input[1].split(',').map((a) => Number(a));

for (let i = 0; i < order.length; i++) {
	if (order[i] < 6)
		final *= Number(yourTickets[i]);
}

console.log(final);
