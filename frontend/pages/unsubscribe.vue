<template>
  <v-container>
    <v-card>
      <v-card-title>Unsubscribe</v-card-title>
      <v-card-text>
        <v-form @submit.prevent="unsubscribe">
          <v-text-field v-model="email" label="Email"></v-text-field>
          <v-btn type="submit">Submit</v-btn>
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
        unsubscribe(){
            this.$axios.get("https://dog.ceo/api/breeds/image/random").then((res) => console.log(res))
        }
    }
}
</script>
