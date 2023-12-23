import { z } from 'zod';
import type { Entry } from './entry';
import type { Unknown } from './utils';

export type Collection = {
	id: string;
	title: string;
	description: string;
	url: string;
};

export type CollectionData = Collection & {
	entries: Entry[];
};

const idSchema = z.string().min(1);
const createOneSchema = z.object({
	title: z.string().min(1),
	description: z.optional(z.string()),
	url: z.optional(z.string())
});
const updateOneSchema = createOneSchema.extend({
	id: idSchema
});

export type CreateOneDto = z.infer<typeof createOneSchema>;
export type UpdateOneDto = z.infer<typeof updateOneSchema>;

export const validators = {
	id: (id: unknown) => idSchema.safeParse(id),
	createOne: (data: Unknown<CreateOneDto>) => createOneSchema.safeParse(data),
	updateOne: (data: Unknown<UpdateOneDto>) => updateOneSchema.safeParse(data)
};

const collectionRepo = {
	validators
};
export default collectionRepo;
