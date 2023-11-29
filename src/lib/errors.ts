import { error } from '@sveltejs/kit';

export const inputError = () =>
	error(400, {
		message: 'wrong input'
	});

export const notFoundError = () =>
	error(404, {
		message: 'not found'
	});
