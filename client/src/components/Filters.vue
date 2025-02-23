<template>
    <div class="mb-4">
      <label class="block mb-2">Date Range</label>
      <input type="date" class="border p-2" v-model="startDate">
      to
      <input type="date" class="border p-2" v-model="end_date">
    </div>
    <div class="mb-4">
      <label class="block mb-2">Sender</label>
      <input type="text" class="border p-2" v-model="sender">
    </div>
    <div class="mb-4">
      <label class="block mb-2">Subject</label>
      <input type="text" class="border p-2" v-model="subject">
    </div>
  </template>
  <script>
  export default {
    data() {
      return {
        startDate: '',
        end_date: '',
        sender: '',
        subject: '',
      }
    },
  }
  </script>