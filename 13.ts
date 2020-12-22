import * as files from 'fs';
import { off } from 'process';

let input = files.readFileSync('13.txt').toString().split('\n');

let buses = input[1].split(',').sort();
buses = buses.slice(0, buses.indexOf('x'));

let times = buses.map((a) => [Number(a) - Number(input[0]) % Number(a), a]).sort((a, b) => Number(a[0]) - Number(b[0]));

console.log('1st Part: ', Number(times[0][0]) * Number(times[0][1]));


buses = input[1].split(',');

function gcd(a:number, b:number) {
	var t = 0;
	a < b && (t = b, b = a, a = t); // swap them if a < b
	t = a % b;
	return t ? gcd(b, t) : b;
}

function lcm(a:number, b:number) {
	return a / gcd(a, b) * b;
}

let i = Number(buses[0]);
let check = 0;
let delta = Number(buses[0]);
while (i < Infinity) {
	check = 0;
	for (let j = 1; j < buses.length; j++) {
		if (buses[j] !== 'x') {
			if (!checkTime(Number(buses[j]), j, i)) {
				break;
			} else {
				delta = lcm(delta, Number(buses[j]));
			}
		}
		check++;
	}
	if (check === buses.length - 1) {
		break;
	}
	i += delta;
}

function checkTime(bus: number, offset: number, time: number) {
	return (time + offset) % bus === 0;
}

console.log('2nd Part: ', i);

