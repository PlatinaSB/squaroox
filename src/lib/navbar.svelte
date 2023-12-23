<script>
	import {
		Navbar,
		NavBrand,
		NavLi,
		NavUl,
		NavHamburger
	} from 'flowbite-svelte';
	import icon from "$lib/assets/favicon.jpg"
	
</script>

<Navbar let:NavContainer color="gray">
	<NavContainer class="border px-5 py-2 rounded-lg bg-white dark:bg-gray-600">
		<NavBrand href="/">
			<img src={icon} class="me-3 h-6 sm:h-9" alt="Squaroox Logo" />
			<span class="self-center whitespace-nowrap text-xl font-semibold">Squaroox</span>
		</NavBrand>
		<NavHamburger />
		<NavUl>
			<NavLi href="/">Home</NavLi>
			<NavLi href="/about">About</NavLi>
			<NavLi href="/tools">Tools</NavLi>
			<NavLi href="/contacts">Contact</NavLi>
		</NavUl>
	</NavContainer>
</Navbar>