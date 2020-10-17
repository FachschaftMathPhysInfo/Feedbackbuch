<template>
  <v-list-item>
    <v-card :elevation="highlighted ? 10 : 2" class="mx-max" style="width: 100vw; margin: 12px">
      <div class="pt-2 pl-2">
        <v-chip v-if="comment.references" label
          ><v-icon class="mr-1">mdi-reply</v-icon>
          <a class="" v-if="comment.references>=0"  v-on:click="jumpToComment(comment.references)" :href="'#' + comment.references" :style=" $vuetify.theme.dark ? 'color:white;' : '' "
            >Kommentar #{{ comment.references }}</a
          >
          <span v-else>gelöschter Kommentar</span>
        </v-chip
        >
      </div>
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
          <span
            ><v-chip :style="highlighted ? {fontWeight: 'bold'} : ''">Kommentar #{{ comment.id }}</v-chip></span
          >
          <span>
            
              <v-dialog v-if="admin" v-model="dialog" persistent max-width="290">
                <template v-slot:activator="{ on, attrs }">
                  <v-btn text v-bind="attrs" v-on="on"
                    ><v-icon>mdi-delete</v-icon> Delete</v-btn
                  ></template
                >
                <v-card>
                  <v-card-title class="headline">
                    Kommentar löschen?
                  </v-card-title>
                  <v-card-text>
                    Möchtest du den Kommentar wirklich unwiderruflich löschen?</v-card-text
                  >
                  <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="primary darken-1" text @click="dialog = false">
                      Nein
                    </v-btn>
                    <v-btn color="green darken-1" text @click="deleteComment">
                      Ja
                    </v-btn>
                  </v-card-actions>
                </v-card>
              </v-dialog>
            
            <v-btn text v-on:click="reply"
              ><v-icon>mdi-reply</v-icon> Reply</v-btn
            >
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
    admin: Boolean,
    highlighted: Boolean,
  },
  data: () => ({
    dialog: false,
    renderConfig: {
      // Mermaid config
      mermaid: {
        theme: "dark",
      },
      texmath: {
        engine: require("katex"),
        katexOptions: {
          macros: { "\\RR": "\\mathbb{R}" },
          output: "html", // formula delimiters
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
                references
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
      this.$emit("reply", this.comment.id);
    },
    jumpToComment(referenceID){
      this.$emit("jumpToComment", referenceID);
    },
    deleteComment() {
      console.log(this.comment.id);
      //TODO:delete comment on server
      this.$apollo.mutate({
        mutation: gql`
          mutation($id: Int!) {
            deleteComment(comment: $id) {
              id
              content
              upvotes
              timestamp
              references
            }
          }
        `,
        variables: {
          id: this.comment.id,
        },
      });
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped></style>
