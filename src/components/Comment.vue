<template>
  <v-list-item>
    <v-card class="mx-max" style="width: 100vw; margin: 12px;">
      <!-- <v-card-text class="headline font-weight-bold">{{comment.text}}</v-card-text> -->
      <Editor
          mode="viewer"
          ref="editor"
          hint="Preview"
          :emoji="true"
          :image="false"
          :outline="false"
          :render-config="renderConfig"
          v-model="comment.content"
        />
      <v-card-actions>
        <v-row justify="end">
          <v-btn text><v-icon>mdi-reply</v-icon> Reply</v-btn>
          <v-btn text v-on:click="upvote">
            <v-icon class="mr-1">mdi-thumb-up</v-icon>
            Upvote {{ this.comment.upvotes }}</v-btn
          >
        </v-row>
      </v-card-actions>
    </v-card>
  </v-list-item>
</template>

<script>
import { Editor } from "vuetify-markdown-editor";

export default {
  name: "Comment",
  components: {
    Editor,
  },
  props: {
    comment: Object,
  },
  data: () => ({
    text: "Tom",
    renderConfig: {
        // Mermaid config
        mermaid: {
          theme: "dark",
        },
        texmath: {
          engine: require("katex"),
          katexOptions: { macros: { "\\RR": "\\mathbb{R}" }, output:"mathml", // formula delimiters
          delimiters: [
            {
              left: "$$$",
              right: "$$$",
              options: {
                displayMode: true, // block
              },
            },
            {
              left: "\\[",
              right: "\\]",
              options: {
                displayMode: true, // block
              },
            },
            {
              left: "$$",
              right: "$$",
              options: {
                displayMode: false, // inline
              },
            },
            {
              left: "\\(",
              right: "\\)",
              options: {
                displayMode: false, // inline
              },
            },
          ],
          },
        },
      },
  }),
  methods: {
    upvote() {
      this.comment.upvotes = this.comment.upvotes + 1;
    },
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>
