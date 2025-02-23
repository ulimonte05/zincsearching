<template>
  <div>
    <p class="mb-2 text-gray-200">{{ searchResults.length }} result(s)</p>
    <search-result 
      v-for="email in searchResults" 
      :key="email._id" 
      :emailSubject="email_subject" 
      :emailSender="email['@timestamp']" 
      :emailDate="email.content" 
      :emailSnippet="email_snippet" 
      :emailId="email._id" 
      @viewDetails="viewDetails">
    </search-result>
  </div>
</template>
<script>
import SearchResult from './SearchResult.vue'
export default {
  components: {
    SearchResult,
  },
  props: {
    searchResults: Array,
  },
  methods: {
    viewDetails(emailId) {
      this.$emit('viewDetails', emailId)
    },
  },
}
</script>