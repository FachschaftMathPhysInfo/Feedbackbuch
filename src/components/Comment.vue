<template>
  <v-list-item>
    <v-card class="mx-max" style="width: 100vw; margin: 12px">
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
        <v-row justify="space-between" style="padding: 0px 12px;">
          <span><i>Kommentar #{{comment.id}}</i></span>
          <span>
          <v-btn text v-on:click="reply"><v-icon>mdi-reply</v-icon> Reply</v-btn>
          <v-btn text v-on:click="upvote">
            <v-icon class="mr-1" :color="upvoted ? 'green' : 'gray'"
              >mdi-thumb-up</v-icon
            >
            Upvote {{ this.comment.upvotes }}</v-btn
          >
          </span>
        </v-row>
      </v-card-actions>
    </v-card>
  </v-list-item>
</template>

<script>
import { Editor } from "vuetify-markdown-editor";
import gql from "graphql-tag";

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
        katexOptions: {
          macros: { "\\RR": "\\mathbb{R}" },
          output: "mathml", // formula delimiters
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
    upvoted: false,
  }),
  methods: {
    upvote() {
      if (this.upvoted) return;
      this.$apollo
        .mutate({
          mutation: gql`
            mutation($id: Int!) {
              upvoteComment(comment: $id) {
                id
                content
                upvotes
                timestamp
              }
            }
          `,
          variables: {
            id: this.comment.id,
          },
        })
        .then(() => (this.upvoted = true));
    },
    reply() {
      this.$emit('reply',this.comment.id);
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>
