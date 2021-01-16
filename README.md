# Employee of the month docker action

Praise your best boy or girl employee in your docs with a lovely gif every month.

E.G.:

```md
<!--- employe-of-the-month-action: Start --->
## ✨✨ Employee of the month December ✨✨

![Employee of the month](https://media3.giphy.com/media/fGFL53eiN8OAAPWd2I/giphy.gif)
<!--- employe-of-the-month-action: End --->
```

## Inputs

### `tag-to-search`:

**Required** Tag to search for your employee. Default: `"dog"`

### `title`:

**Required** Tag to search for your employee. Default: `"✨✨ Employee of the month {{.Month}}✨✨"`

### `page`:

**Required** Page to update. Default: `"Home.md"`

## Example usage

```yml
on:
  schedule:
    # * is a special character in YAML so you have to quote this string
    - cron: "* 0 1 * *"
...
- name: Employer Of The Month
  uses: KirioXX/employee-of-the-month-action@v1.0.0
  with:
    tag-to-search: "cat"
    title: "{{.Month}} good kitty of the month"
    page: "Home.md"
  env:
    GIPHY_API_KEY: ${{secrets.GIPHY_API_KEY}}
    GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
...
```
