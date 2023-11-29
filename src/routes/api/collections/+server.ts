import { json } from '@sveltejs/kit';
import type { RequestHandler } from '@sveltejs/kit';
import { inputError } from '$lib/errors';
import { collectionRepo } from '$lib/server/db';

export const GET: RequestHandler = async () => {
	const collections = await collectionRepo.findAll();

	return json(collections);
};

export const POST: RequestHandler = async ({ request }) => {
	const data = await request.json();

	const result = collectionRepo.validators.createOne({
		title: data.title,
		description: data.description,
		url: data.url
	});
	if (!result.success) {
		throw inputError();
	}

	const createdCollection = await collectionRepo.createOne(result.data);

	return json(createdCollection);
};
