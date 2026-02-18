<script setup lang="ts">
import { onMounted, ref } from 'vue';


type Job = {
	id: string;
	title: string;
	url: string;
	description: string;
}

const jobs = ref<Job[]>([])

onMounted(async () => {
	const response = await fetch('http://localhost:8080/getAll')

	jobs.value = await response.json()
})
</script>

<template>
	<main>
		<h1>Jobs</h1>
		<div v-for="job in jobs" :key="'job-list' + job.id">
			<h2>{{ job.title }}</h2>
			<p>{{ job.description }}</p>
			<a :href="job.url" target="_blank">{{ job.url }}</a>
		</div>
	</main>
</template>

<style scoped></style>
