<template>
  <v-container>
    <v-card>
      <v-card-title>Admin</v-card-title>
      <v-card-text>
        <v-form @submit.prevent="sendContent">
          <v-textarea v-model="content" label="Newsletter Content" filled></v-textarea>
          <v-btn type="submit">Send to all subscribers!</v-btn>
        </v-form>
      </v-card-text>
    </v-card>
  </v-container>
</template>

<script>
export default {
    data: () => ({
        content: '',
    }),
    methods:{
        sendContent(){
            this.$axios.get("https://dog.ceo/api/breeds/image/random").then((res) => console.log(res))
        }
    }
}
</script>
