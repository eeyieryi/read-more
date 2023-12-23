import type { NumericRange } from '@sveltejs/kit';

type MyError = {
	status: NumericRange<400, 599>;
	message: string;
};

export const unableToLoadError: MyError = {
	status: 404,
	message: 'unable to load data'
};
