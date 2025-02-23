<template>
    <div class="border p-4">
      <h2 class="text-2xl font-semibold">Email Details</h2>
      <p class="text-gray-500">Content: {{ email.content }}</p>
      <p class="text-gray-500">To: {{ email.to }}</p>
      <p class="text-gray-500">Subject: {{ email_subject }}</p>
      <p class="text-gray-500">Date: {{ email.date }}</p>
      <div v-html="email.body"></div>
    </div>
  </template>
  <script>
  export default {
    props: {
      email: Object,
    },
  }
  </script>