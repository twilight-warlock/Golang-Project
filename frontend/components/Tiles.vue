<template>
  <v-card max-width="500" class="d-flex justify-center align-center">
    <v-row>
      <v-col cols="8">
        <v-card-title>{{ this.topic }}</v-card-title>
      </v-col>
      <v-col cols="4">
        <v-card-action>
          <v-btn>
            <v-icon> mdi-plus </v-icon>
          </v-btn>
        </v-card-action>
      </v-col>
    </v-row>
  </v-card>
</template>

<script>
export default {
  props: ['topic'],
}
</script>

<style></style>
