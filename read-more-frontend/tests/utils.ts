export function getPathname(url: string) {
	return new URL(url).pathname;
}
