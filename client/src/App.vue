<template>
  <div class="flex">
    <div class="w-full">
      <Header/>
      <div class="container mx-auto w-full flex flex-col items-center justify-center p-4">
        <SearchBar @search="performSearch" />
        <Filters v-if="showFilters" />
        <div v-if="searchResults === null" class="text-center text-white">
          No results.
        </div>
        <SearchResultsList v-else :searchResults="searchResults" @viewDetails="viewEmailDetails" />
      </div>
    </div>
    <!-- <div class="w-1/2">
      <EmailDetail v-if="selectedEmail" :email="selectedEmail" />
    </div> -->
  </div>
</template>

<script>
import Header from './components/Header.vue'
import SearchBar from './components/SearchBar.vue'
import Filters from './components/Filters.vue'
import SearchResultsList from './components/SearchResultsList.vue'
import EmailDetail from './components/EmailDetail.vue'
import axios from 'axios'

export default {
  components: {
    Header,
    SearchBar,
    Filters,
    SearchResultsList,
    EmailDetail,
  },
  data() {
    return {
      searchResults: [],
      selectedEmail: null,
      showFilters: false,
    }
  },
  methods: {
    performSearch(query) {
      if (!query) {
        query = "email"
      }
      axios.post('http://localhost:8080/emails/search', {
        query 
      }).then(response => {
        this.searchResults = response.data
      }).catch(error => {
        console.error(error)
      })
    },
    viewEmailDetails(emailId) {
      axios.get(`http://localhost:3000/emails/${emailId}`).then(response => {
        this.selectedEmail = response.data
      }).catch(error => {
        console.error(error)
      })
    },
  },
}
</script>