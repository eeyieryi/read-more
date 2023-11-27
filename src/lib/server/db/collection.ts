import { z } from 'zod';
import prisma, { type Collection, type CollectionData } from './prisma';
import type { Unknown } from '$lib/utils';

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

export async function findAll(): Promise<Collection[]> {
	// TODO: Handle exceptions for Prisma
	return await prisma.collection.findMany();
}

export async function findOneById(id: string): Promise<CollectionData | null> {
	return await prisma.collection.findUnique({
		where: { id },
		include: { entries: true }
	});
}

export async function createOne(
	createOneDto: CreateOneDto
): Promise<Collection> {
	return await prisma.collection.create({
		data: {
			title: createOneDto.title,
			description: createOneDto.description ?? '',
			url: createOneDto.url ?? ''
		}
	});
}

export async function updateOne(
	updateOneDto: UpdateOneDto
): Promise<Collection> {
	return await prisma.collection.update({
		where: {
			id: updateOneDto.id
		},
		data: {
			title: updateOneDto.title,
			description: updateOneDto.description,
			url: updateOneDto.url
		}
	});
}

export async function deleteOneById(id: string): Promise<Collection> {
	return await prisma.collection.delete({
		where: { id }
	});
}
