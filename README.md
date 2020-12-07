# Employee of the month docker action

Praise your best boy employee in your docs with a love gif every month.

E.G.:

```md
<!--- employe-of-the-month-action: Start --->
## ✨✨ Employee of the month ✨✨

![Employee of the month](https://media3.giphy.com/media/fGFL53eiN8OAAPWd2I/giphy.gif)
<!--- employe-of-the-month-action: End --->
```

## Inputs

### `GIPHY_API_KEY`

**Required** Your giphy api key (keep it secret).

### `tag-to-search`:

**Required** Tag to search for your employee. Default: `"dog"`

### `title`:

**Required** Tag to search for your employee. Default: `"✨✨ Employee of the month ✨✨"`

## Example usage

```yml
uses: actions/employe-of-the-month-action@v1
with:
  GIPHY_API_KEY: ${{secrets.GIPHY_API_KEY}}
```
