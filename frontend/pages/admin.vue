<template>
  <v-container>
    <v-card color="blue darken-2">
      <v-card-title>Admin</v-card-title>
      <v-card-text>
        <v-form @submit.prevent="sendContent">
          <v-textarea v-model="content" label="Newsletter Content" color="white" filled></v-textarea>
          <!-- <v-checkbox v-model="isHTML" label="Is HTML" value="isHTML" color="white"></v-checkbox> -->
          <v-btn color="blue darken-1" depressed type="submit">Send to all subscribers!</v-btn>
        </v-form>
        <i>Supports HTML!</i>

        <br />

        <h1 class="mt-10">Subscribers </h1>
        <p>...</p>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script>
export default {
    data: () => ({
        content: '',
        // isHTML: false,
    }),
    methods:{
        sendContent(){
            this.$axios.get("https://dog.ceo/api/breeds/image/random").then((res) => console.log(res))
        }
    }
}
</script>
