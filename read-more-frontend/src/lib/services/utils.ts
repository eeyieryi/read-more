export type Unsuccessful = {
	success: false;
};

export type Successful = {
	success: true;
};

export type SuccessfulWithObject<T, K extends string> = Successful & {
	[P in K]: T;
};

export type SuccessfulWithId = Successful & {
	id: string;
};

export type ReturnTypeBase = {
	success: boolean;
};

export type ReturnTypeBaseWithId = ReturnTypeBase & {
	id: string;
};

export type CreateOneReturnType<T, K extends string> = Promise<
	SuccessfulWithObject<T, K> | Unsuccessful
>;

export type UpdateOneReturnType<T, K extends string> = Promise<
	SuccessfulWithObject<T, K> | Unsuccessful
>;
