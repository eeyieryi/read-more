<script lang="ts">
	export let src: string;

	let audio: HTMLMediaElement;

	function handleKeyDown(e: KeyboardEvent) {
		if (e.key === ' ') {
			e.preventDefault();
			if (audio.paused) {
				audio.play();
			} else {
				audio.pause();
			}
		} else if (e.key === 'ArrowLeft') {
			e.preventDefault();
			audio.currentTime -= 5;
		} else if (e.key === 'ArrowRight') {
			e.preventDefault();
			audio.currentTime += 5;
		}
	}
</script>

<svelte:window on:keydown="{handleKeyDown}" />

<audio controls bind:this="{audio}">
	<source src="{src}" />
</audio>
