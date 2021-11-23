<template>
  <v-container>
    <v-card color="teal">
      <v-card-title>Subscribe</v-card-title>
      <v-card-text>
        <v-form @submit.prevent="subscribe">
          <v-text-field v-model="email" label="Email" color="teal lighten-4"></v-text-field>
          <v-btn type="submit" color="teal lighten-1">Subscribe</v-btn>
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
    methods:{
        subscribe(){
            this.$axios.get("https://dog.ceo/api/breeds/image/random").then((res) => console.log(res))
        }
    }
}
</script>
