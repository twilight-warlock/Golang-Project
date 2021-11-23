<template>
  <v-app dark>
    <v-main>
      <v-app-bar>
        <v-app-bar-title
          ><nuxt-link to="/" style="text-decoration: none; color: inherit"
            >GoLetter</nuxt-link
          ></v-app-bar-title
        >
        <v-spacer />
        <nuxt-link to="/subscribe" style="text-decoration: none; color: inherit"
          ><v-btn class="p-1" text>Subscribe</v-btn></nuxt-link
        >
        <nuxt-link
          to="/unsubscribe"
          style="text-decoration: none; color: inherit"
          ><v-btn class="p-1" text>Unsubscribe</v-btn></nuxt-link
        >
        <nuxt-link to="/admin" style="text-decoration: none; color: inherit"
          ><v-btn class="p-1" text>Admin</v-btn></nuxt-link
        >
      </v-app-bar>
      <v-container>
        <Nuxt />
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
export default {
  data() {
    return {
      clipped: false,
      drawer: false,
      fixed: false,
      items: [
        {
          icon: 'mdi-apps',
          title: 'Welcome',
          to: '/',
        },
        {
          icon: 'mdi-chart-bubble',
          title: 'Inspire',
          to: '/inspire',
        },
      ],
      miniVariant: false,
      right: true,
      rightDrawer: false,
      title: 'Vuetify.js',
    }
  },
}
</script>
