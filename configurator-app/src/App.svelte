<script>
	import Header from './components/Header.svelte'
	import DeviceList from './components/DeviceList.svelte'
	import DeviceEdit from './components/DeviceEdit.svelte'
	import SystemStatus from './components/SystemStatus.svelte'

	var editDeviceId = null
	var editOpened = false

	function handleDeviceClick(ev) {
		if (ev.detail) {
			editOpened = true
			editDeviceId = ev.detail.id || null
		} else {
			editOpened = false
		}
	}
</script>

<style>
	:global(body) {
		background: #f6f6f6;
	}

	main {
		max-width: 720px;
		margin: 4rem auto;
	}
</style>


<main>
	<Header />
	<SystemStatus />
	{#if editOpened}
		<DeviceEdit deviceId={editDeviceId} />
	{/if}
	<DeviceList on:deviceClicked={handleDeviceClick} />
</main>
