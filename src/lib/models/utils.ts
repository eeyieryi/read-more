export type Unknown<T> = {
	[P in keyof T]: unknown;
};
