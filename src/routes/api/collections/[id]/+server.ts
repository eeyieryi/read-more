import { json } from '@sveltejs/kit';
import type { RequestHandler } from '@sveltejs/kit';
import { inputError, notFoundError } from '$lib/errors';
import { collectionRepo } from '$lib/server/db';

export const GET: RequestHandler = async ({ params }) => {
	const result = collectionRepo.validators.id(params.id);
	if (!result.success) {
		throw inputError();
	}

	const collection = await collectionRepo.findOneById(result.data);
	if (!collection) {
		throw notFoundError();
	}

	return json(collection);
};

export const DELETE: RequestHandler = async ({ params }) => {
	const result = collectionRepo.validators.id(params.id);
	if (!result.success) {
		throw inputError();
	}

	const deletedCollection = await collectionRepo.deleteOneById(result.data);

	return json(deletedCollection);
};

export const PUT: RequestHandler = async ({ params, request }) => {
	const data = await request.json();

	const result = collectionRepo.validators.updateOne({
		id: params.id,
		title: data.title,
		description: data.description,
		url: data.url
	});
	if (!result.success) {
		throw inputError();
	}

	const updatedCollection = await collectionRepo.updateOne(result.data);

	return json(updatedCollection);
};
