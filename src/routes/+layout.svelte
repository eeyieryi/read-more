<script lang="ts">
	import { page } from '$app/stores';
	import '../app.pcss';

	let isShowing = true;

	const toggleVisibility = () => (isShowing = !isShowing);
</script>

<div class="flex h-screen flex-col overflow-hidden bg-gray-50">
	<nav class="flex flex-row border-b-4 border-b-gray-500">
		{#if isShowing}
			<a
				class:active="{$page.route.id === '/collections' ||
					$page.route.id === '/collections/[id]'}"
				class="link text-center"
				href="/collections">Collections</a>
			<a
				class:active="{$page.route.id === '/blank'}"
				class="link text-center"
				href="/blank">Blank</a>
			<a
				class:active="{$page.route.id === '/collections/new'}"
				class="link ml-auto"
				href="/collections/new">New Collection</a>
			<a
				class:active="{$page.route.id === '/entries/new'}"
				class="link"
				href="/entries/new">New Entry</a>
			<button class="btn" on:click|preventDefault="{toggleVisibility}"
				>-</button>
		{:else}
			<button
				class="btn ml-auto"
				on:click|preventDefault="{toggleVisibility}">+</button>
		{/if}
	</nav>
	<slot />
</div>

<style lang="postcss">
	.link,
	.btn {
		@apply block px-4 py-2 text-gray-900;
	}
	.active {
		@apply bg-gray-500 text-white;
	}
</style>
