import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
	name: 'truncate'
})
export class TruncatePipe implements PipeTransform {

	transform(value: number): number {
		return parseFloat(value.toFixed(3));
	}

}
