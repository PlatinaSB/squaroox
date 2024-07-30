<script>
	import { Button, Dropdown, DropdownItem, Label, Input } from 'flowbite-svelte';
	import { ChevronDownSolid } from 'flowbite-svelte-icons';
	let gb1 = 'Pilih bangun datar';
	let dropdownOpen = false;
	import Navbar from '$lib/navbar.svelte';

	let val = '';
	let val2 = '';
	let result = '';
	const Persegi = async () => {
		try {
			const response = await fetch('/api/Persegi', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					value: val
				})
			});

			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}

			const json = await response.json();
			result = json.result;
		} catch (error) {
			console.error('Error:', error);
		}
	};

	const Lingkaran = async () => {
		try {
			const response = await fetch('http://localhost:8000/api/Lingkaran', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					value: val
				})
			});

			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}

			const json = await response.json();
			result = json.result;
		} catch (error) {
			console.error('Error:', error);
		}
	};

	const Segitiga = async () => {
		try {
			const response = await fetch('http://localhost:8000/api/Segitiga', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					value: val,
					value2: val2
				})
			});

			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}

			const json = await response.json();
			result = json.result;
		} catch (error) {
			console.error('Error:', error);
		}
	};

	const Persegi_Panjang = async () => {
		try {
			const response = await fetch('http://localhost:8000/api/Persegi_Panjang', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json'
				},
				body: JSON.stringify({
					value: val,
					value2: val2
				})
			});

			if (!response.ok) {
				throw new Error(`HTTP error! Status: ${response.status}`);
			}

			const json = await response.json();
			result = json.result;
		} catch (error) {
			console.error('Error:', error);
		}
	};
</script>

<Navbar />

<div class="px-5 mx-5 bg-gray-400 min-h-80 rounded-lg">
	<div class="pt-5 items-center">
		<Button>{gb1}<ChevronDownSolid class="w-3 h-3 ms-2 text-white dark:text-white" /></Button>
		<Dropdown class="w-44 p-3 space-y-3 text-sm" bind:open={dropdownOpen}>
			<DropdownItem on:click={() => ((dropdownOpen = false), (gb1 = 'Persegi'))}
				>Persegi</DropdownItem
			>
			<DropdownItem on:click={() => ((dropdownOpen = false), (gb1 = 'Lingkaran'))}
				>Lingkaran</DropdownItem
			>
			<DropdownItem on:click={() => ((dropdownOpen = false), (gb1 = 'Segitiga'))}
				>Segitiga</DropdownItem
			>
			<DropdownItem on:click={() => ((dropdownOpen = false), (gb1 = 'Persegi_Panjang'))}
				>Persegi Panjang</DropdownItem
			>
		</Dropdown>
	</div>
	<div class="pt-5 pb-2">
		{#if gb1 == 'Persegi'}
			<Label for="default-input" class="block mb-2">Masukan ukuran Persegi</Label>
			<Input bind:value={val} placeholder="Masukan ukuran Persegi" />
			<Button on:click={Persegi} color="light" class="my-3">hitung</Button>
			<Label for="default-input" class="block mb-2">hasil</Label>
			<Input id="disabled-input" class="mb-6" disabled placeholder="hasil" value={result} />
		{:else if gb1 == 'Lingkaran'}
			<Label for="default-input" class="block mb-2">Masukan radius Lingkaran</Label>
			<Input bind:value={val} id="Lingkaran-radius" placeholder="Masukan radius Lingkaran" />
			<Button on:click={Lingkaran} color="light" class="my-3">hitung</Button>
			<Label for="default-input" class="block mb-2">hasil</Label>
			<Input id="disabled-input" class="mb-6" disabled placeholder="hasil" value={result} />
		{:else if gb1 == 'Segitiga'}
			<Label for="default-input" class="block mb-2">Masukkan alas</Label>
			<Input bind:value={val} id="Segitiga-alas" placeholder="Masukan ukuran alas" />
			<Label for="default-input" class="block mb-2">Masukkan tinggi</Label>
			<Input bind:value={val2} id="Segitiga-tinggi" placeholder="Masukan ukuran tinggi" />
			<Button on:click={Segitiga} or="light" class="my-3">hitung</Button>
			<Label for="default-input" class="block mb-2">hasil</Label>
			<Input id="disabled-input" class="mb-6" disabled placeholder="hasil" value={result} />
		{:else if gb1 == 'Persegi_Panjang'}
			<Label for="default-input" class="block mb-2">Masukan ukuran panjang</Label>
			<Input bind:value={val} id="Persegi_Panjang-panjang" placeholder="Masukan ukuran panjang" />
			<Label for="default-input" class="block mb-2">Masukan ukuran lebar</Label>
			<Input bind:value={val2} id="Persegi_Panjang-lebar" placeholder="Masukan ukuran lebar" />
			<Button on:click={Persegi_Panjang} color="light" class="my-3">hitung</Button>
			<Label for="default-input" class="block mb-2">hasil</Label>
			<Input id="disabled-input" class="mb-6" disabled placeholder="hasil" value={result} />
		{/if}
	</div>
</div>
