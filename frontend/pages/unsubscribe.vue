<template>
  <v-container>
    <v-card color="brown">
      <v-card-title>Unsubscribe</v-card-title>
      <v-card-text>
        <v-form @submit.prevent="unsubscribe">
          <v-text-field v-model="email" label="Email"></v-text-field>
          <v-btn type="submit" color="brown lighten-1">Unsubscribe</v-btn>
        </v-form>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script>
export default {
  data: () => ({
    email: '',
  }),
  methods: {
    unsubscribe() {
      this.$axios.post('http://localhost:8080/unsubscribe', {
        email: this.email,
      })
    },
  },
}
</script>
