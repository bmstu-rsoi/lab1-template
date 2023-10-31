using Microsoft.EntityFrameworkCore;
using RSOI_Lab_01.Entities;
using RSOI_Lab_01.Interfaces;

namespace RSOI_Lab_01.Repositories
{
    public class PersonsRepository : IPersonsRepository
    {
        readonly PersonsDbContext _context;

        public PersonsRepository(PersonsDbContext context)
        {
            _context = context;
        }

        public async Task CreateAsync(Person person)
        {
            await _context.Persons.AddAsync(person).ConfigureAwait(false);
        }

        public async Task DeleteAsync(int id)
        {
            var item = await GetAsync(id).ConfigureAwait(false);

            if (item is not null)
                _context.Persons.Remove(item);
        }

        public async Task<IEnumerable<Person>> GetAllAsync()
        {
            return await _context.Persons.ToListAsync().ConfigureAwait(false);
        }

        public async Task<Person?> GetAsync(int id)
        {
            return await _context.Persons.FirstOrDefaultAsync(e => e.Id.Equals(id)).ConfigureAwait(false);
        }

        public async Task SaveAsync()
        {
            await _context.SaveChangesAsync().ConfigureAwait(false);
        }

        public void Update(Person person)
        {
            _context.Persons.Update(person);
        }
    }
}
