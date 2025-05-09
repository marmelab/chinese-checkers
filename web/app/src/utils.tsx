/**
 * Merge multiple class names to one full class name.
 * @param className Class names.
 */
export function classes(...className: (string | null | undefined | false)[]): string {
	return className.filter((className) => !!className).join(" ");
}
