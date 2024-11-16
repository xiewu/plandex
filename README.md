<h1 align="center">
 <a href="https://plandex.ai">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="images/plandex-logo-dark.png"/>
    <source media="(prefers-color-scheme: light)" srcset="images/plandex-logo-light.png"/>
    <img width="400" src="images/plandex-logo-dark-bg.png"/>
 </a>
 <br />
</h1>
<br />

<div align="center">

<p align="center">
  <!-- Call to Action Links -->
  <a href="#install">
    <b>30-Second Install</b>
  </a>
   · 
  <a href="https://plandex.ai">
    <b>Website</b>
  </a>
   · 
  <a href="https://docs.plandex.ai/">
    <b>Docs</b>
  </a>
   · 
  <a href="#more-examples-">
    <b>Examples</b>
  </a>
   · 
  <a href="https://docs.plandex.ai/hosting/self-hosting">
    <b>Self-Hosting</b>
  </a>
</p>

<br>

[![Discord](https://img.shields.io/discord/1214825831973785600.svg?style=flat&logo=discord&label=Discord&refresh=1)](https://discord.gg/plandex-ai)
[![GitHub Repo stars](https://img.shields.io/github/stars/plandex-ai/plandex?style=social)](https://github.com/plandex-ai/plandex)
[![Twitter Follow](https://img.shields.io/twitter/follow/PlandexAI?style=social)](https://twitter.com/PlandexAI)

</div>

<p align="center">
  <!-- Badges -->
<a href="https://github.com/plandex-ai/plandex/pulls"><img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg" alt="PRs Welcome" /></a> <a href="https://github.com/plandex-ai/plandex/releases?q=cli"><img src="https://img.shields.io/github/v/release/plandex-ai/plandex?filter=cli*" alt="Release" /></a>
<a href="https://github.com/plandex-ai/plandex/releases?q=server"><img src="https://img.shields.io/github/v/release/plandex-ai/plandex?filter=server*" alt="Release" /></a>

  <!-- <a href="https://github.com/your_username/your_project/issues">
    <img src="https://img.shields.io/github/issues-closed/your_username/your_project.svg" alt="Issues Closed" />
  </a> -->

</p>

<br />

<div align="center">
<a href="https://trendshift.io/repositories/8994" target="_blank"><img src="https://trendshift.io/api/badge/repositories/8994" alt="plandex-ai%2Fplandex | Trendshift" style="width: 250px; height: 55px;" width="250" height="55"/></a>
</div>

<br>

<h3 align="center">AI driven development in your terminal.<br/>Build entire features and apps with a robust workflow.</h3>

<br/>
<br/>

<!-- Vimeo link is nicer on mobile than embedded video... downside is it navigates to vimeo in same tab (no way to add target=_blank) -->
<!-- https://github.com/plandex-ai/plandex/assets/545350/c2ee3bcd-1512-493f-bdd5-e3a4ca534a36 -->

<a href="https://player.vimeo.com/video/926634577">
  <img src="images/plandex-intro-vimeo.png" alt="Plandex intro video" width="100%"/>
</a>

<br/>
<br/>

## More examples  🎥

<h4>👉  <a href="https://www.youtube.com/watch?v=0ULjQx25S_Y">Building Pong in C/OpenGL with GPT-4o and Plandex</a></h4>

<h4>👉  <a href="https://www.youtube.com/watch?v=rnlepfh7TN4">Fixing a tricky real-world bug in 5 minutes with Claude Opus 3 and Plandex</a></h4>

<br/>

## Learn more about Plandex  🧐

- [Overview](#overview-)
- [Install](#install)
- [Hosting options](#hosting-options-)
- [Get started](#get-started-)
- [Docs](https://docs.plandex.ai/)
- [Build complex software](#build-complex-software-with-llms-)
- [Why Plandex?](#why-plandex-)
- [Discussion and discord](#discussion-and-discord-)
- [Contributors](#contributors-)
<br/>

## Overview  📚

<p>Churn through your backlog, work with unfamiliar technologies, get unstuck, and <strong>spend less time on the boring stuff.</strong></p>

<p>Plandex is a <strong>reliable and developer-friendly</strong> AI coding agent in your terminal. It can plan out and complete <strong>large tasks</strong> that span many files and steps.</p>
 
<p>Designed for <strong>real-world use-cases</strong>, Plandex can help you build a new app quickly, add new features to an existing codebase, write tests and scripts, understand code, fix bugs, and automatically debug failing commands (like tests, typecheckers, linters, etc.). </p>

<br/>

## Install  📥

```bash
curl -sL https://plandex.ai/install.sh | bash
```

**Note:** Windows is supported via [WSL](https://learn.microsoft.com/en-us/windows/wsl/install). Plandex only works correctly on Windows in the WSL shell. It doesn't work in the Windows CMD prompt or PowerShell.

[More installation options.](https://docs.plandex.ai/install)

<br/>

## Hosting options ⚖️

| # | Option  | Description |
|---|---------|--------------------------------|
| 1 | **Plandex Cloud (Integrated Models)** | No separate accounts or API keys are required. This is the quickest way to get started. If you choose this option, skip ahead to the [Get Started](#get-started-) section below. |
| 2 | **Plandex Cloud (BYO API Key)** | You'll need accounts and API keys for [OpenRouter.ai](https://openrouter.ai) and [OpenAI](https://platform.openai.com) to get started with the default models. |
| 3 | **Self-hosted** | First, follow the [self-hosting guide](./hosting/self-hosting.md) to set up your own Plandex server. You'll also need accounts and API keys for [OpenRouter.ai](https://openrouter.ai) and [OpenAI](https://platform.openai.com) to get started with the default models. |

If you're going with option 2 or 3 above, you'll need to set the `OPENROUTER_API_KEY` and `OPENAI_API_KEY` environment variables before continuing:

```bash
export OPENROUTER_API_KEY=...
export OPENAI_API_KEY=...
```

## Get started  🚀

Now `cd` into your **project's directory.** Make a new directory first with `mkdir your-project-dir` if you're starting on a new project.

```bash
cd your-project-dir
```


Then **start your first plan** with `plandex new`.

```bash
plandex new
```

☁️ If you're using Plandex Cloud, you'll be prompted at this point to start a trial.

Load any relevant files, directories, directory layouts, urls, or images **into the LLM's context** with `plandex load`.

```bash
plandex load some-file.ts another-file.ts
plandex load src/components -r # load a whole directory
plandex load src --tree # load a directory layout (file names only)
plandex load src/**/*.ts # load files matching a glob pattern
plandex load https://raw.githubusercontent.com/plandex-ai/plandex/main/README.md # load the text content of a url
plandex load images/mockup.png # load an image
```


Now **send your prompt.** You can pass it in as a file:

```bash
plandex tell -f prompt.txt
```


Write it in vim:

```bash
plandex tell # tell with no arguments opens vim so you can write your prompt there
```


Or pass it inline (use enter for line breaks):

```bash
plandex tell "add a new line chart showing the number of foobars over time to components/charts.tsx"
```

Plandex will make a plan for your task and then implement that plan in code. **The changes won't yet be applied to your project files.** Instead, they'll accumulate in Plandex's sandbox.

**Note**: if you're not quite ready to give Plandex a task yet and want to ask questions or chat a bit first, you can use `plandex chat` instead of `plandex tell`. It works the same way, but it makes Plandex respond conversationally and prevents it from making any changes yet. Once you're ready, you can use `plandex tell` to go ahead with the implementation.

```bash
plandex chat "is it clear from the context how to add a new line chart?"
```

To learn about reviewing changes, iterating on the plan, and applying changes to your project, **[continue with the full quickstart.](https://docs.plandex.ai/quick-start#review-the-changes)**

<br/>

## Docs  🛠️

### [👉  Full documentation.](https://docs.plandex.ai/)


<br/>

## Build complex software with LLMs  🌟

⚡️  Changes are accumulated in a protected sandbox so that you can review them before automatically applying them to your project files. Built-in version control allows you to easily go backwards and try a different approach. Branches allow you to try multiple approaches and compare the results.

📑  Manage context efficiently in the terminal. Easily add files or entire directories to context, and keep them updated automatically as you work so that models always have the latest state of your project.

🧠  By default, Plandex uses a mix of Anthropic models (via OpenRouter.ai) and OpenAI models. You can also use it with a wide range of other models from any OpenAI-compatible provider.

✅  Plandex supports Mac, Linux, FreeBSD, and Windows. It runs from a single binary with no dependencies.

<br/>

## Why Plandex?  🤔

🏗️  Go beyond autocomplete to build complex functionality with AI.<br>
🚫  Stop the mouse-centered, copy-pasting madness of coding with ChatGPT or Claude.<br>
⚡️  Ensure the model always has the latest versions of files in context.<br>
🪙  Retain granular control over what's in the model's context and how many tokens you're using.<br>
⏪  Rewind, iterate, and retry as needed until you get your prompt just right.<br>
❤️‍🩹  Automatically debug failing commands (like tests, typecheckers, linters, etc.).<br>
🌱  Explore multiple approaches with branches.<br>
🔀  Run tasks in the background or work on multiple tasks in parallel.<br>
🎛️  Try different models and temperatures, then compare results.<br>

<br/>

## Discussion and discord  💬

Please feel free to give your feedback, ask questions, report a bug, or just hang out:

- [Discord](https://discord.gg/plandex-ai)
- [Discussions](https://github.com/plandex-ai/plandex/discussions)
- [Issues](https://github.com/plandex-ai/plandex/issues)

<br/>

## Contributors  👥

⭐️  Please star, fork, explore, and contribute to Plandex. There's a lot of work to do and so much that can be improved.

[Here's an overview on setting up a development environment.](https://docs.plandex.ai/development)


