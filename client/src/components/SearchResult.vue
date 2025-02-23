<template>
    <div class="border p-4 mb-4 hover:shadow-md">
      <h2 class="text-xl font-semibold">{{ emailSubject }}</h2>
      <p class="text-gray-500">From: {{ emailSender }}</p>
      <p class="text-gray-500">Date: {{ emailDate }}</p>
      <p class="font-bold text-blue-200 ">id: {{ emailId }}</p>
    </div>
  </template>
  <script>
  export default {
    props: {
      emailSubject: String,
      emailSender: String,
      emailDate: String,
      emailSnippet: String,
      emailId: String,
    },
    methods: {
      viewDetails() {
        this.$emit('viewDetails', this.emailId)
      },
    },
  }
  </script>