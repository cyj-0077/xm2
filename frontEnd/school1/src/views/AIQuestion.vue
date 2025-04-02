<!-- xm/项目/frontEnd/school1/src/views/AIQuestion.vue -->
<template>
    <div>
        <input v-model="question" placeholder="请输入问题">
        <button @click="askQuestion">提问</button>
        <p v-if="answer">{{ answer }}</p>
    </div>
</template>

<script setup>
import { ref } from 'vue';
import { aiQuestionAnswer } from '@/utils/request';

const question = ref('');
const answer = ref('');

async function askQuestion() {
    try {
        answer.value = await aiQuestionAnswer(question.value);
    } catch (error) {
        console.error('问答出错:', error);
    }
}
</script>