<script setup lang="ts">
import { onMounted, ref } from 'vue';

type Job = {
	id: string;
	title: string;
	url: string;
	description: string;
}

const jobs = ref<Job[]>([])

// Accordion ids
const openIds = ref<string[]>([])

function toggle(toggleId: string) {
	if (openIds.value.includes(toggleId)) {
		openIds.value = openIds.value.filter(id => id !== toggleId)
	} else {
		openIds.value.push(toggleId)
	}
}

const newJobUrl = defineModel("")
const error = ref('')

async function addNewJob() {
	if (newJobUrl.value === undefined) return

	try {
		const response = await fetch(`http://localhost:8080/add?jobUrl=${newJobUrl.value}`, { method: "post" })

		const newJob = await response.json()

		jobs.value.push(newJob)

	} catch {
		error.value = "Could not add job, sorry!"
	}
}

onMounted(async () => {
	const response = await fetch('http://localhost:8080/getAll')
	jobs.value = await response.json()
})
</script>

<template>
	<main class="min-h-screen bg-gray-800 text-white">
		<div class="mx-auto w-full max-w-4xl px-4 sm:px-6 lg:px-8 py-10">
			<h1 class="text-3xl font-semibold tracking-tight mb-8">Jobs</h1>
			<div class="flex">
				<input type="text" placeholder="New job url" v-model="newJobUrl"
					class="w-full bg-white rounded-xl px-2 placeholder:text-neutral-800 text-black" />
				<button class="inline-flex items-center justify-center rounded-lg ml-4 bg-violet-600 px-4 py-2 text-sm font-medium text-white
         hover:bg-violet-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition-colors"
					@click="addNewJob()">
					Add new job
				</button>
			</div>

			<article v-for="job in jobs" :key="'job-list' + job.id"
				class="rounded-xl border border-neutral-200 p-6 shadow-sm my-4">
				<div class="flex justify-between" @click="toggle(job.id)">
					<h2>{{ job.title }}</h2>
					<span class="transition-transform duration-200 cursor-pointer"
						:class="{ 'rotate-180': openIds.includes(job.id) }">
						▼
					</span>
				</div>
				<a class="text-blue-400 hover:underline" :href="job.url" target="_blank">link</a>
				<section v-show="openIds.includes(job.id)">
					<p>{{ job.description }}</p>
				</section>
			</article>
		</div>
	</main>
</template>

<style scoped></style>
