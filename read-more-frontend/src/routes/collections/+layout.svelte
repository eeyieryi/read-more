<script lang="ts">
	import type { LayoutData } from './$types';
	import { page } from '$app/stores';

	export let data: LayoutData;
</script>

<div class="flex h-full flex-row overflow-auto">
	<div
		class="scrollable flex min-w-[200px] flex-col border-r-4 border-gray-500">
		{#each data.collections as collection (collection.id)}
			<a
				class="block w-full py-3 text-center"
				class:active="{$page.params.id === collection.id}"
				href="/collections/{collection.id}">{collection.title}</a>
		{/each}
	</div>
	<slot />
</div>

<style lang="postcss">
	.active {
		@apply bg-gray-500 text-white;
	}
</style>
