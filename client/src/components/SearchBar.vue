<template>
    <div class="flex items-center mb-4">
      <input type="text" placeholder="Search emails" class="border p-2 text-white rounded-l" v-model="searchQuery" @keypress.enter="search">
      <button class="text-black p-2 rounded-r" @click="search">Search</button>
    </div>
  </template>
  <script>
  export default {
    data() {
      return {
        searchQuery: '',
      }
    },
    methods: {
      search() {
        this.$emit('search', this.searchQuery)
      },
    },
  }
  </script>