/* Resize the table to the minimal size that contains all the elements,
 * but with the invariant of a USED/BUCKETS ratio near to <= 1 */
int dictResize(dict *d)
{
	int minimal;

	if (!dict_can_resize || dictIsRehashing(d)) 
		return DICT_ERR;
	minimal = d->ht[0].used;
	if (minimal < DICT_HT_INITIAL_SIZE)
		minimal = DICT_HT_INITIAL_SIZE;
	return dictExpand(d, minimal);
}

/* Expand or create the hash table */
int dictExpand(dict *d, unsigned long size)
{
	dictht n; /* the new hash table */
	unsigned long realsize = _dictNextPower(size);

	/* the size is invalid if it is smaller than the number of
	 * elements already inside the hash table */
	if (dictIsRehashing(d) || d->ht[0].used > size)
		return DICT_ERR;

	/* Rehashing to the same table size is not useful. */
	if (realsize == d->ht[0].size) return DICT_ERR;

	/* Allocate the new hash table and initialize all pointers to NULL */
	n.size = realsize;
	n.sizemask = realsize-1;
	n.table = zcalloc(realsize*sizeof(dictEntry*));
	n.used = 0;

	/* Is this the first initialization? If so it's not really a rehashing
	 * we just set the first hash table so that it can accept keys. */
	if (d->ht[0].table == NULL) {
		d->ht[0] = n;
		return DICT_OK;
	}

	/* Prepare a second hash table for incremental rehashing */
	d->ht[1] = n;
	d->rehashidx = 0;
	return DICT_OK;
}

/* Performs N steps of incremental rehashing. Returns 1 if there are still
 * keys to move from the old to the new hash table, otherwise 0 is returned.
 *
 * Note that a rehashing step consists in moving a bucket (that may have more
 * than one key as we use chaining) from the old to the new hash table, however
 * since part of the hash table may be composed of empty spaces, it is not
 * guaranteed that this function will rehash even a single bucket, since it
 * will visit at max N*10 empty buckets in total, otherwise the amount of
 * work it does would be unbound and the function may block for a long time. */
int dictRehash(dict *d, int n) {
	int empty_visits = n*10; /* Max number of empty buckets to visit. */
	if (!dictIsRehashing(d)) return 0;

	while(n-- && d->ht[0].used != 0) {
		dictEntry *de, *nextde;

		/* Note that rehashidx can't overflow as we are sure there are more
		 * elements because ht[0].used != 0 */
		assert(d->ht[0].size > (unsigned long)d->rehashidx);
		while(d->ht[0].table[d->rehashidx] == NULL) {
			d->rehashidx++;
			if (--empty_visits == 0) return 1;
		}
		de = d->ht[0].table[d->rehashidx];
		/* Move all the keys in this bucket from the old to the new hash HT */
		while(de) {
			unsigned int h;

			nextde = de->next;
			/* Get the index in the new hash table */
			h = dictHashKey(d, de->key) & d->ht[1].sizemask;
			de->next = d->ht[1].table[h];
			d->ht[1].table[h] = de;
			d->ht[0].used--;
			d->ht[1].used++;
			de = nextde;
		}
		d->ht[0].table[d->rehashidx] = NULL;
		d->rehashidx++;
	}

	/* Check if we already rehashed the whole table... */
	if (d->ht[0].used == 0) {
		zfree(d->ht[0].table);
		d->ht[0] = d->ht[1];
		_dictReset(&d->ht[1]);
		d->rehashidx = -1;
		return 0;
	}

	/* More to rehash... */
	return 1;
}
