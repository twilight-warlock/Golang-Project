<template>
  <v-container>
    <v-card color="blue darken-2">
      <v-card-title>Admin</v-card-title>
      <v-card-text>
        <v-form @submit.prevent="sendContent">
          <v-text-field
            v-model="subject"
            label="Subject"
            color="white"
            filled
          ></v-text-field>
          <v-textarea
            v-model="content"
            label="Newsletter Content"
            color="white"
            filled
          ></v-textarea>
          <i>Supports HTML!</i> <br />
          <!-- <v-checkbox v-model="isHTML" label="Is HTML" value="isHTML" color="white"></v-checkbox> -->
          <v-btn color="blue darken-1" class="mt-2" depressed type="submit"
            >Send to all subscribers!</v-btn
          >
        </v-form>

        <br />

        <h1 class="mt-10 mb-5">Subscribers</h1>
        <p class="p-2" v-for="email in emails" :key="email">{{ email }}</p>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script>
export default {
  data: () => ({
    subject: '',
    content: '',
    emails: [],
    // isHTML: false,
  }),
  async mounted() {
    const { data } = await this.$axios.get('http://localhost:8080/subscribers')
    this.emails = data
  },
  methods: {
    sendContent() {
      this.$axios.post('http://localhost:8080/send_newsletter', {
        title: this.subject,
        body: this.content,
      })
    },
  },
}
</script>
