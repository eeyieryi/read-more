import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types';
import apiService from '$lib/services/api.service';
import { unableToLoadError } from '$lib/errors';

export const load: PageLoad = async ({ fetch, params: { id } }) => {
	const res = await apiService.entry.findOne(fetch, id);
	if (!res.success) {
		const err = unableToLoadError;
		error(err.status, err.message);
	}
	const resTwo = await apiService.collection.findAll(fetch);
	if (!resTwo.success) {
		const err = unableToLoadError;
		error(err.status, err.message);
	}
	return {
		entryData: res.entryData,
		collections: resTwo.data
	};
};
